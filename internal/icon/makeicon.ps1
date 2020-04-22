# makeicon icon_file variable_nane package_name go_file
# e.g. makeicon icon.ico IconDataArray icon newicon.go

# https://stackoverflow.com/a/61357962/6396652
Start-Process cmd -ArgumentList "/c TYPE $($args[0]) | 2goarray $($args[1]) $($args[2]) > $($args[3])"
