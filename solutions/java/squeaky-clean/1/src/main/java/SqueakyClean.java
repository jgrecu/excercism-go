
class SqueakyClean {
    static String clean(String identifier) {
        var result = new StringBuilder();
        var capital = false;
        for (int i = 0; i < identifier.length(); i++) {
            if (identifier.charAt(i) == ' ') {
                result.append('_');
            } else if (identifier.charAt(i) == '-') {
                capital = true;
                continue;
            } else if (capital) {
                result.append(Character.toUpperCase(identifier.charAt(i)));
                capital = false;
            } else if (identifier.charAt(i) == '4') {
                result.append('a');
            } else if (identifier.charAt(i) == '3') {
                result.append('e');
            } else if (identifier.charAt(i) == '0') {
                result.append('o');
            } else if (identifier.charAt(i) == '1') {
                result.append('l');
            } else if (identifier.charAt(i) == '7') {
                result.append('t');
            } else if (!Character.isLetter(identifier.charAt(i))) {
                continue;
            } else {
                result.append(identifier.charAt(i));
            }
        }
        return result.toString();
    }
}
