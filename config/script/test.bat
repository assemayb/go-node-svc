@echo off

set commandsFile=test-commands.txt

for /f "tokens=*" %%a in (%commandsFile%) do (
  start cmd /k %%a
)