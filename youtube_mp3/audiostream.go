package youtube_mp3

import (
	"io"
	"log"
	"os/exec"
)

type AudioStream struct {
	downloadInvocation  *exec.Cmd
	transcodeInvocation *exec.Cmd
	output              io.ReadCloser
}

func NewAudioStream(videoID string) (stream *AudioStream, err error) {
	log.Printf("DEBUG: Getting video %s", videoID)
	stream = &AudioStream{}
	stream.downloadInvocation = downloadVideoInvocation(videoID)
	/*webm, err := stream.downloadInvocation.StdoutPipe()
	if err != nil {
		return
	}*/
	stream.output, err = stream.downloadInvocation.StdoutPipe()
	if err != nil {
		return
	}
	/*stream.transcodeInvocation = transcodeMP3Invocation(webm)
	stream.output, err = stream.transcodeInvocation.StdoutPipe()
	if err != nil {
		return
	}*/
	err = stream.downloadInvocation.Start()
	if err != nil {
		return
	}
	/*err = stream.transcodeInvocation.Start()
	if err != nil {
		return
	}*/
	return
}

func (stream *AudioStream) Read(buffer []byte) (int, error) {
	return stream.output.Read(buffer)
}

func (stream *AudioStream) Close() error {
	stream.downloadInvocation.Process.Kill()
	//stream.transcodeInvocation.Process.Kill()
	//err := stream.transcodeInvocation.Wait()
	err := stream.downloadInvocation.Wait()
	return err
}

func downloadVideoInvocation(videoID string) *exec.Cmd {
	invocation := exec.Command("youtube-dl", videoID, "-o", "-", "-f", "bestaudio") // youtube-dl <videoID> -o - -f bestaudio
	invocation.Stderr = log.Writer()
	log.Printf("DEBUG: youtube-dl: %s\n", invocation.String())

	return invocation
}

/*func transcodeMP3Invocation(webm io.ReadCloser) *exec.Cmd {
	invocation := exec.Command("ffmpeg", "-i", "pipe:", "-f", "mp3", "-") // ffmpeg -i pipe: -f mp3 -
	invocation.Stderr = log.Writer()
	invocation.Stdin = webm
	log.Printf("DEBUG: ffmpeg: %s\n", invocation.String())

	return invocation
}*/
