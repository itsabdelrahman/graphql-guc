const API = 'https://m.guc.edu.eg/StudentServices.asmx';

export const urls = {
  LOGIN: API.concat('/Login'),
  ATTENDANCE: API.concat('/GetAttendance'),
  COURSEWORK: API.concat('/GetCourseWork'),
  EXAMS_SCHEDULE: API.concat('/GetExamsSchedule'),
  SCHEDULE: API.concat('/GetSchedule'),
  TRANSCRIPT: API.concat('/GetTranscript'),
};

export const config = {
  CLIENT_VERSION: '1.3',
  APP_OS: '0',
  OS_VERSION: '6.0.1',
};
