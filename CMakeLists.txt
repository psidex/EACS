cmake_minimum_required(VERSION 3.13)
project(E_APO_Config_Switcher C)
set(CMAKE_C_STANDARD 99)

add_executable(E_APO_Config_Switcher main.c helpers.c helpers.h)

# Move icon over to binary dir so it is seen by the bin
file(COPY ${CMAKE_CURRENT_SOURCE_DIR}/icon.ico DESTINATION ${CMAKE_CURRENT_BINARY_DIR})
