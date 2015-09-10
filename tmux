#!/bin/bash
set -e

session=modern-web-apps-in-go

function sourceDevEnv {
	tmux send-keys -t $1 "source dev.env" C-m
}

if ! tmux ls | grep -q "$session"; then
	tmux new-session -d -s $session

	tmux rename-window vim
	tmux split-window -v

	sourceDevEnv $session:1.1
	tmux send-keys -t $session:1.1 "vim" C-m

	sourceDevEnv $session:1.2
	tmux resize-pane -t $session:1.2 -D 20

	tmux new-window -n "convey"
	sourceDevEnv $session:2
	tmux send-keys -t $session:2 "goconvey" C-m
	open http://localhost:8080
fi

tmux select-window -t $session:1
tmux select-pane -L
exec tmux attach-session -d -t $session
