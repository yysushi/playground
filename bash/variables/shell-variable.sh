#!/bin/bash

HOGE1=FUGA1
export HOGE2=FUGA2

TMPFILE=$(mktemp)
cat << EOF > "$TMPFILE"
echo \$HOGE1
echo \$HOGE2
EOF

. "$TMPFILE"
# source "$TMPFILE"
bash "$TMPFILE"
