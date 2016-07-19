# dql-client
The standalone DQL Client

This is a prototype to demonstrate that GOLang will enable us to build a working DQL client that works exactly like the MySQL client.

It handles commands delimited by a semi-colin ';'.
Like MySQL, it can handle commands that span across multiple lines and it keeps a history of the commands you've run (during this session).

It can also handle copy pasting a large chunk of text into the terminal, try it out with the text chunk below.

```
Single line Command;
Multiline
command;
Multiple line
with incomplete command; Starting a new command midway
and not finishing it
```
