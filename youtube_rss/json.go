package youtube_rss

import (
	"io"
	"log"
	"os/exec"
)

type JSONStream struct {
	invocation *exec.Cmd
	output     io.ReadCloser
}

func NewJSONStream(path string) (stream *JSONStream, err error) {
	log.Printf("DEBUG: Getting playlist %s", path)
	stream = &JSONStream{}
	stream.invocation = downloadJSONInvocation(path)

	stream.output, err = stream.invocation.StdoutPipe()
	if err != nil {
		return
	}

	err = stream.invocation.Start()
	if err != nil {
		return
	}

	return
}

func (stream *JSONStream) Read(buffer []byte) (int, error) {
	return stream.output.Read(buffer)
}

func (stream *JSONStream) Close() error {
	stream.invocation.Process.Kill()
	err := stream.invocation.Wait()
	return err
}

func downloadJSONInvocation(path string) *exec.Cmd {
	invocation := exec.Command("youtube-dl", "youtube.com/"+path, "-J") // youtube-dl <addr> -J
	invocation.Stderr = log.Writer()
	log.Printf("DEBUG: youtube-dl: %s\n", invocation.String())

	return invocation
}
