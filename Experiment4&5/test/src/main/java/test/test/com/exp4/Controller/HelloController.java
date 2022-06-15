package test.test.com.exp4.Controller;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
@Controller
public class HelloController {
    @RequestMapping("/hello")
    public String hello(){
        return "hello";
    }

    @RequestMapping("/index")
    public String index(){
        return "index";
    }
}
