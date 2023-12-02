# Thoughts
In this file I'll write down what I thought about the puzzle and my own solution.

## Day 1: Trebuchet?!
Part 1 was quite simple, I used Regex my beloved to find all the numbers and get the first and the last one.
Nice little puzzle to get started with but also so easy it gave me a false sense of security.

Part 2 was a different story. In the end the code is quite simple but to get to the point where I understood the logic it took quite a while. If you haven't done it yet, you do have to account for overlapping matches, this was **not** in the instructions. This means a solution with a simple strings Replacer is wrong :(  

A friend of mine introduced me to the concept of sliding/moving windows which was new to me as I never studied computer science, using that logic I came to a pretty elegant solution. Some guys on the Subreddit were replacing words with numerals padded with part of the word. While I suppose that solution also gets you the answer, I do prefer mine.

## Day 2: Cube Conundrum
Part 1 took me the most time because that's where I did most of the work on parsing the input to something usable. I think I might be becoming a bit too reliant on using maps to move data around, but it's just so easy. Can't say I'm aware of negative impact of using maps either, maybe they use more memory, but luckily this is not a microcontroller and I can have all the memory I want >:3c (spoken like web app developers all over the world)

Part 2 was simple because I already had everything neatly parsed and had code blocks ready to look through game data