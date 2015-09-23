#!/bin/bash
set -e

session=modern-web-apps-in-go

function sourceDevEnv {
	tmux send-keys -t $1 "source dev.env" C-m
}

if ! tmux ls | grep -q "$session"; then
	tmux new-session -d -s $session

	sourceDevEnv $session:1.1
	tmux send-keys -t $session:1.1 "vim" C-m

    tmux split-window -h
    tmux send-keys -t $session:1.2 "curl -i localhost:8888"

	tmux split-window -v

	sourceDevEnv $session:1.3
	tmux resize-pane -t $session:1.3 -D 20
    tmux send-keys -t $session:1.3 "go run cmd/server/main.go" C-m
fi

tmux select-window -t $session:1
tmux select-pane -L
exec tmux attach-session -d -t $session
