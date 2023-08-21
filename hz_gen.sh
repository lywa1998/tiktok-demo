#!/bin/sh
if [ "$1" = "init" ]; then
    hz new -mod tiktok_demo
fi

hz update -idl idl/common.thrift

hz update -idl idl/user.thrift
hz update -idl idl/comment.thrift
hz update -idl idl/feed.thrift
hz update -idl idl/favorite.thrift
hz update -idl idl/message.thrift
hz update -idl idl/publish.thrift
hz update -idl idl/relation.thrift
