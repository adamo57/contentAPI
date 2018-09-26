<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Api;

/**
 * Post Service.
 */
class PostClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * GetPost takes in certain query params and returns a Post.
     * @param \Api\PostSlug $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetPost(\Api\PostSlug $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/api.Post/GetPost',
        $argument,
        ['\Api\PostObject', 'decode'],
        $metadata, $options);
    }

}
