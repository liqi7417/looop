package main

func main() {
    var s_str string = "323"
    s_str_list := []string{s_str}
    print(s_str, "\n")
    print(s_str_list[0], "\n")
    s_str_list[0] = "456"
    print(s_str, "\n")
    print(s_str_list[0], "\n")
}
