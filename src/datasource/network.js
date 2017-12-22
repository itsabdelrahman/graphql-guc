import axios from 'axios';
import { urls, config } from './constants';

const request = url => credentials =>
  axios.post(url, {
    clientVersion: config.CLIENT_VERSION,
    app_os: config.APP_OS,
    os_version: config.OS_VERSION,
    username: credentials.username,
    password: credentials.password,
  });

export const requestLogin = request(urls.LOGIN);

export const requestAttendance = request(urls.ATTENDANCE);

export const requestCoursework = request(urls.COURSEWORK);

export const requestExams = request(urls.EXAMS_SCHEDULE);

export const requestSchedule = request(urls.SCHEDULE);

export const requestTranscript = request(urls.TRANSCRIPT);
