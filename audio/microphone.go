package audio

import (
	"os"
	"time"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"github.com/gordonklaus/portaudio"
	"github.com/wrk-grp/errnie"
)

type Microphone struct {
	sampleRate int
	threshold  float32
	buffer     []float32
	silence    int
	recording  bool
}

func NewMicrophone() *Microphone {
	errnie.Handles(portaudio.Initialize())
	return &Microphone{44100, 0.01, make([]float32, 0), 0, false}
}

func (microphone *Microphone) Record() {
	microphone.recording = true

	stream, err := portaudio.OpenDefaultStream(
		1, 0, float64(microphone.sampleRate), len(microphone.buffer),
		func(in []float32) {
			if !microphone.recording {
				return
			}

			for _, i := range in {
				if i > microphone.threshold {
					microphone.silence = 0
				}

				if i < microphone.threshold && i > -microphone.threshold {
					microphone.silence++
				}

				microphone.buffer = append(microphone.buffer, i)
			}

			if microphone.silence >= microphone.sampleRate {
				microphone.recording = false
				microphone.silence = 0
			}
		},
	)

	errnie.Handles(err)
	stream.Start()
	defer stream.Close()

	for microphone.recording {
		time.Sleep(100 * time.Millisecond)
	}

	microphone.writeWav()
}

func (microphone *Microphone) Close() {
	errnie.Handles(portaudio.Terminate())
}

func (microphone *Microphone) writeWav() {
	wavFile, err := os.Create("out.wav")
	errnie.Handles(err)
	defer wavFile.Close()

	wavEncoder := wav.NewEncoder(wavFile, microphone.sampleRate, 16, 1, 1)
	intBuffer := make([]int, len(microphone.buffer))

	for i, f := range microphone.buffer {
		intBuffer[i] = int(f * 32767)
	}

	audioBuffer := &audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: 1,
			SampleRate:  microphone.sampleRate,
		},
		Data:           intBuffer,
		SourceBitDepth: 16,
	}

	err = wavEncoder.Write(audioBuffer)
	errnie.Handles(err)
	wavEncoder.Close()
	microphone.buffer = make([]float32, 0)
}
