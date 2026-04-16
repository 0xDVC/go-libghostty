cmake_minimum_required(VERSION 3.19)
project(go-libghostty LANGUAGES C)

include(FetchContent)
FetchContent_Declare(ghostty
    GIT_REPOSITORY https://github.com/ghostty-org/ghostty.git
    GIT_TAG 9e080c5a403475dcbee93c40eeb22cf6f92121f4
)
FetchContent_MakeAvailable(ghostty)
