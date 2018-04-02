import R from 'ramda';
import axios from 'axios';
import { urls, config } from './constants';
import {
  areMockCredentials,
  mockLoginRequest,
  mockAttendanceRequest,
  mockCourseworkRequest,
  mockExamsRequest,
  mockScheduleRequest,
  mockTranscriptRequest,
} from './mock';

const request = url => credentials =>
  axios.post(url, {
    clientVersion: config.CLIENT_VERSION,
    app_os: config.APP_OS,
    os_version: config.OS_VERSION,
    username: credentials.username,
    password: credentials.password,
  });

export const requestLogin = R.ifElse(
  areMockCredentials,
  mockLoginRequest,
  request(urls.LOGIN),
);

export const requestAttendance = R.ifElse(
  areMockCredentials,
  mockAttendanceRequest,
  request(urls.ATTENDANCE),
);

export const requestCoursework = R.ifElse(
  areMockCredentials,
  mockCourseworkRequest,
  request(urls.COURSEWORK),
);

export const requestExams = R.ifElse(
  areMockCredentials,
  mockExamsRequest,
  request(urls.EXAMS_SCHEDULE),
);

export const requestSchedule = R.ifElse(
  areMockCredentials,
  mockScheduleRequest,
  request(urls.SCHEDULE),
);

export const requestTranscript = R.ifElse(
  areMockCredentials,
  mockTranscriptRequest,
  request(urls.TRANSCRIPT),
);
