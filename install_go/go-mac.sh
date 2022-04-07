# brew install golang 
#! /bin/sh
shell_type="$(echo $SHELL)"


arr=(${shell_type//"/"/ })
echo $arr
# read -ra arr <<< "$shell_type"
# len=${#arr[@]}

# echo "Current shell =>" ${arr[$len-1]}
# curr_shell=${arr[$len-1]}
# curr_shell="zsh"
# shell_rc=".${curr_shell}rc"
# p="$HOME/$shell_rc"
# echo $p

# echo "Installing Go -----"

# brew install golang



