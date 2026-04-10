cmake_minimum_required(VERSION 3.19)
project(go-libghostty LANGUAGES C)

include(FetchContent)
FetchContent_Declare(ghostty
    GIT_REPOSITORY https://github.com/ghostty-org/ghostty.git
    GIT_TAG c34901dddbc36b63f33bbb1a47d62f6911584d65
)
FetchContent_MakeAvailable(ghostty)
