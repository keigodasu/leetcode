func interpret(command string) string {
    var output bytes.Buffer
    temp := []rune{}
    dic := map[string]string{
        "G": "G",
        "()":  "o",
        "(al)": "al",
    }
    
    for _, v := range command {
        temp = append(temp, v)
        if val, ok := dic[string(temp)]; ok {
            output.WriteString(val) 
            temp = []rune{}
        }
    }
    
    return output.String()
}

