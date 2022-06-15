package test.test.com.exp4.Controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.serializer.StringRedisSerializer;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;

@Component
public class RedisBean {

    @Autowired
    private RedisTemplate redisTemplate;

    public static RedisTemplate redis;

    @PostConstruct
    public void getRedisTemplate(){
        redis=this.redisTemplate;
        redis.setKeySerializer(new StringRedisSerializer());
        redis.setValueSerializer(new StringRedisSerializer());
    }

}
