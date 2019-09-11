package msgproc

import "testing"

func TestImageCommand_IsRun(t *testing.T) {
	type args struct {
		update BotUpdate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "gbr no param", args: args{update: TestUpdate{message: "gbr "}}, want: false},
		{name: "gbr with param", args: args{update: TestUpdate{message: "gbr mega"}}, want: true},
		{name: "gbr in the middle", args: args{update: TestUpdate{message: "adadas gbr asdada"}}, want: false},
		{name: "gbr with double space", args: args{update: TestUpdate{message: "gbr    "}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imageCommand := ImageCommand{}
			if got := imageCommand.IsExecuted(tt.args.update); got != tt.want {
				t.Errorf("IsRun() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestUpdate struct {
	message string
}

func (t TestUpdate) Message() string {
	return t.message
}
