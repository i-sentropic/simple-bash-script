#!/bin/bash


# Generate random duration and calculate number of iterations
duration=$((RANDOM % 4 + 4))
iterations=$((duration))

cat file_contents.txt

# Random words list
WORDS=("apple" "banana" "cherry" "date" "elderberry" "fig" "grape" "honeydew" "kiwi" "lemon" "mango" "nectarine" "orange" "papaya" "quince" "raspberry" "strawberry" "tangerine" "ugli" "vanilla" "watermelon" "xigua" "yam" "zucchini")

echo "duration: $((duration))"

for ((i=0; i<iterations; i++)); do

    # Generate a random number of words (e.g., 3 to 7 words)
    num_words=$((RANDOM % 5 + 3))
    output=""

    for ((j = 0; j < num_words; j++)); do
        rand_index=$((RANDOM % ${#WORDS[@]}))
        output+="${WORDS[$rand_index]} "
    done

    # Alternate between stdout and stderr
    if (( i % 2 == 0 )); then
        echo "STDERR: $(date '+%H:%M:%S') - $output" >&2
    else
        echo "STDOUT: $(date '+%H:%M:%S') - $output"
    fi

    sleep 1

done




