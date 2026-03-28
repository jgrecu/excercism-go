public class LogLevels {
    
    public static String message(String logLine) {
        return logLine.substring(logLine.indexOf(':') + 1).strip();
    }

    public static String logLevel(String logLine) {
        return logLine.substring(logLine.indexOf('[') + 1, logLine.indexOf(']')).strip().toLowerCase();
    }

    public static String reformat(String logLine) {
        String msg = logLine.substring(logLine.indexOf(':') + 1).strip();
        String level = logLine.substring(logLine.indexOf('[') + 1, logLine.indexOf(']')).strip().toLowerCase();
        return msg + " (" + level + ")";
    }
}
