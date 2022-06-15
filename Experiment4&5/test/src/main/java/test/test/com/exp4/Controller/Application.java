package test.test.com.exp4.Controller;

import org.json.JSONObject;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;


@SpringBootApplication
@RestController
public class Application {

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
    @GetMapping("/like")
    public String like(@RequestParam(value = "pageName") String pageName) {
        RedisTemplate redisTemplate = test.test.com.exp4.Controller.RedisBean.redis;
        Object jsonstrobj = redisTemplate.opsForValue().get(pageName);
        JSONObject node = new JSONObject(jsonstrobj.toString());
        node.put("likes",node.getInt("likes")+1);
        redisTemplate.opsForValue().set(pageName,node.toString());
        return "Done";
    }
    @GetMapping("/unlike")
    public String unlike(@RequestParam(value = "pageName") String pageName) {
        RedisTemplate redisTemplate = test.test.com.exp4.Controller.RedisBean.redis;
        Object jsonstrobj = redisTemplate.opsForValue().get(pageName);
        JSONObject node = new JSONObject(jsonstrobj.toString());
        node.put("likes",node.getInt("likes")-1);
        redisTemplate.opsForValue().set(pageName,node.toString());
        return "Done";
    }
    @GetMapping("/SeeLike")
    public int SeeLike(@RequestParam(value = "pageName") String pageName) {
        RedisTemplate redisTemplate = test.test.com.exp4.Controller.RedisBean.redis;
        Object jsonstrobj = redisTemplate.opsForValue().get(pageName);
        if(jsonstrobj == null){
            JSONObject newJson = new JSONObject();
            newJson.put("likes",0);
            redisTemplate.opsForValue().set(pageName,newJson.toString());
            return 0;
        }
        JSONObject node = new JSONObject(jsonstrobj.toString());
        System.out.println(node.get("likes"));
        return node.getInt("likes");
    }
}

