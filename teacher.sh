#!/bin/bash

# step 1: Isolate the number of the key interview.
# Try several likely locations for the streets folder so the script works in
# different training/audit layouts.
STREETS_DIR=""
if [ -d "./streets" ]; then
	STREETS_DIR="./streets"
elif [ -d "./mystery/streets" ]; then
	STREETS_DIR="./mystery/streets"
elif [ -d "./the-final-cl-test/mystery/streets" ]; then
	STREETS_DIR="./the-final-cl-test/mystery/streets"
fi

KEY_INTERVIEW_NUMBER=""
if [ -n "$STREETS_DIR" ]; then
	KEY_INTERVIEW_NUMBER=$(grep -h "SEE INTERVIEW" "$STREETS_DIR"/* 2>/dev/null | grep -oE "[0-9]+" | head -n 1)
fi

# step 2: Print the newly created environment variable.
echo "$KEY_INTERVIEW_NUMBER"

# step 3: Print what the interview contains.
# Try several likely Interviews directories/filenames.
INTERVIEWS_DIR=""
if [ -d "./Interviews" ]; then
	INTERVIEWS_DIR="./Interviews"
elif [ -d "./interviews" ]; then
	INTERVIEWS_DIR="./interviews"
elif [ -d "./mystery/interviews" ]; then
	INTERVIEWS_DIR="./mystery/interviews"
elif [ -d "./the-final-cl-test/mystery/interviews" ]; then
	INTERVIEWS_DIR="./the-final-cl-test/mystery/interviews"
fi

if [ -n "$KEY_INTERVIEW_NUMBER" ] && [ -n "$INTERVIEWS_DIR" ]; then
	# Try capitalized and lowercase filename variants
	if [ -f "$INTERVIEWS_DIR/Interview-$KEY_INTERVIEW_NUMBER" ]; then
		cat "$INTERVIEWS_DIR/Interview-$KEY_INTERVIEW_NUMBER"
	elif [ -f "$INTERVIEWS_DIR/interview-$KEY_INTERVIEW_NUMBER" ]; then
		cat "$INTERVIEWS_DIR/interview-$KEY_INTERVIEW_NUMBER"
	fi
fi

# step 4: Print the content of the environment variable MAIN_SUSPECT.
echo "$MAIN_SUSPECT"