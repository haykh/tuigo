package debug

type Debugger struct {
	enabled bool
	msg     string
}

func New() Debugger {
	return Debugger{msg: "", enabled: false}
}

func (d *Debugger) Log(msg string) {
	if d.enabled {
		d.msg = msg
	}
}

func (d *Debugger) Enable() {
	d.enabled = true
}

func (d *Debugger) Disable() {
	d.enabled = false
}

func (d Debugger) Enabled() bool {
	return d.enabled
}

func (d Debugger) Get() string {
	return d.msg
}
