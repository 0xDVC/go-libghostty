cmake_minimum_required(VERSION 3.19)
project(go-libghostty LANGUAGES C)

include(FetchContent)
FetchContent_Declare(ghostty
    GIT_REPOSITORY https://github.com/ghostty-org/ghostty.git
    GIT_TAG 2ed382a15566b267c32fae440b065f7844b15bfb
)
FetchContent_MakeAvailable(ghostty)
