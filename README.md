1. Please share a sample of some code you’ve written or contributed to, that is a good reflection of your programming skill.  Github link(s) would be ideal.  If you're up for writing some fresh code, please take a shot at this programming challenge: https://github.com/replicatedcom/core-challenge-troubleshoot-analyzers

2. Please share a sample of some technical writing you’ve done, particularly if you have something published -- a blog post, design document, block comments from code, etc...
> See: [Service Asset and Configuration Management](./ServiceAssetandConfigurationManagement-161119-2143.pdf)

3. Please share a sample of analysis - for example, a bug report, threat modeling, performance-tuning analysis, or a war story about a security incident or a production outage.
> I had a hard time finding a recent example. I haven't had to formalize any of the above for a couple years.

1. What is an example of something that you learned that was a struggle for you?
> When I see what I perceive to be an injustice to someone, I have the tendacy to want to lash out in their defense. This often does more harm than good. I've had to learn how to check my emotions in these moments and not follow my first instict. I still get angry about the injustice, but I'm able to see it more as a warning sign that something is not quite right and should explore the situation more carefully.

1. Who is someone who has mentored you, and what did you learn from them?
> DKG was a jerk. He was crass, irreverent and stubborn. He wasn't always right, but when he was he'd make sure you knew it. I hated him right up until I joined another team and all of a sudden I realized he had made it possible for me to work with literally anyone. The miserable experience he provided was so much more challenging intellectually and emotionally than anything I have faced since. He taught me how to check my ego at the door, patiently wade through the noise, find what looked like value and finally test that to ensure success. 

6. Who is someone you have mentored, and what did you learn from them?
> David is my youngest brother, we're 6 years apart. He lived near me throughout his college years and naturally we became quite close. I helped him get his first job and always advised him when asked. At some point I had come to a ceiling in my job at the time. I felt guilt about leaving the area to find more oppurtunity. Not to be dramatic, but I felt like I was abondoning him. I brought this up one night and he said, "Um yeah... you should go for it". Almost like he didn't understand why I would feel guilty.
> ...

1. What qualities do you most admire in other software engineers?
> Candidness, eloquence, humor and empathy (understanding the users needs)

8. Describe some of the more interesting differences between two programming languages you know.
> This may be clichè, but Golang's lack of classes is terribly refreshing. In my work up to this point, I've always been on the infrastructure/sysadmin side. My Python code has always been much more procedural than object oriented. When I met a problem that did require classes, I found it difficult to immediately jump into creating classes and often needed to brush up on certain topics even in order to build somthing simple. In Golang, this has become simpler because you're always going to use a struct. Additionally, having functions defined outside of a the struct, allows for more code reuse without the need for the OOP concepts, such as inheritence. Finally, it's nice having a `switch` statment instead strong arming a `dictionary` to act like one. 

9.  Describe a new feature that was added to a programming language/library/framework/tool that you use, and how you decided whether to adopt and use this new feature or not.
> Python's `f-strings`- I avoided using the `str.format()` and prefered the original _hack_ of using the operator `%`. It always felt more concise and dare I say _pythonic_. Enter `f-strings`, this was the most natural transition for me. I liked how it matched the pattern established by `r-strings`, `u-strings` and `b-strings`.

10.  Consider the following scenario:  You are the newest member of a team, working on a new service that consists of several code repositories and their build/test/deployment configurations.  It’s late in the day, and a coworker asks you to code review their most recent pull request, and it’s a big one with lots of changes.  You pull their branch to test it out, and the code just doesn’t work.  You let the coworker know, and they tell you “Sorry, it’s working fine on my computer, and I need to head out, can we revisit tomorrow?”.  The next day, you do a ‘pull’ on this branch, and you don’t see any new commits, but the code is now working fine.  What are some of the things you’d do next, and why?
> I would do a fresh checkout in a new directory of the their PR and see if the problem is occurs again. This would be to see if there might be issues with my working copy of the repo. Essentially trying to rule out _user_ error or a misunderstanding of how a test environment is spun up.
> Next I would look at the history of the other code repositories and see if there were any changes that occured between the time the code didn't work and now. Assuming there are I would dig in to determine how those changes might have affected the PR code. Something else could have changed that allowed the new code to work. 
> Finally, if those two things didn't reveal why it hadn't worked the day before. I'd see if anyone else on the team had expirienced anything similar. Or just make a note to check it out more if it happens again.