package test.test.com.exp4.Controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
public class ReturnHtml {
    @RequestMapping("/favorite")
    public String main() {
        return "./favorite";
    }
}
