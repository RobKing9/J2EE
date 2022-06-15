package test.test.com.exp4.Controller;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;

import java.util.ArrayList;
import java.util.List;

@Controller
public class MainController {
    @GetMapping("/main")
    public String main(Model model){
        List<Content> list = new ArrayList<>();
        String[] strings = {
                "Golang",
                "Linux",
                "Docker",
                "MySQL",
                "Redis"
        };
        for(int i=0;i<strings.length;++i){
            Content content = new Content();
            content.setGoal(strings[i]);
            list.add(content);
        }
        model.addAttribute("list",list);
        return "main";
    }
}
class Content{
    private String content;

    public String getContent() {
        return content;
    }

    public void setGoal(String content) {
        this.content = content;
    }
}