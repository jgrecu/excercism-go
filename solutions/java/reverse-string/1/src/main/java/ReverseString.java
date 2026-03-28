class ReverseString {

    String reverse(String inputString) {
        StringBuilder result = new StringBuilder();
        for (int i = 0; i < inputString.length(); i++) {
            var c = inputString.charAt(i);
            result.insert(0, c);
        }
        return result.toString();
    }
  
}
