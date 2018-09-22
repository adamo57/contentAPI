<?php

require dirname(__FILE__).'/vendor/autoload.php';

function test()
{
    $client = new ContentAPI/PostClient('localhost:8000',[
    'credentials' => Grpc/ChannelCredentials::createInsecure(),
    ]);

    return ""
}

echo test;