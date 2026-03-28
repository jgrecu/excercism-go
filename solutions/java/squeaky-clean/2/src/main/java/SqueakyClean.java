
class SqueakyClean {
    static String clean(String identifier) {
        var result = new StringBuilder();
        var doUpper = false;
        for (int i = 0; i < identifier.length(); i++) {
            char c = identifier.charAt(i);
            if (c == ' ') {
                result.append('_');
            } else if (c == '-') {
                doUpper = true;
            } else if (c == '4') {
                result.append('a');
            } else if (c == '3') {
                result.append('e');
            } else if (c == '0') {
                result.append('o');
            } else if (c == '1') {
                result.append('l');
            } else if (c == '7') {
                result.append('t');
            } else if (doUpper) {
                result.append(Character.toUpperCase(c));
                doUpper = false;
            } else if (!Character.isLetter(c)) {
                continue;
            } else {
                result.append(c);
            }
        }
        return result.toString();
    }
}
