import mockLoginResponse from './mock/login.json';
import mockAttendanceResponse from './mock/attendance.json';
import mockCourseworkResponse from './mock/coursework.json';
import mockExamsResponse from './mock/exams.json';
import mockScheduleResponse from './mock/schedule.json';
import mockTranscriptResponse from './mock/transcript.json';

const mockCredentials = {
  USERNAME: 'john.doe',
  PASSWORD: '123456',
};

export const areMockCredentials = credentials =>
  credentials.username === mockCredentials.USERNAME &&
  credentials.password === mockCredentials.PASSWORD;

export const mockLoginRequest = async () => mockLoginResponse;

export const mockAttendanceRequest = async () => mockAttendanceResponse;

export const mockCourseworkRequest = async () => mockCourseworkResponse;

export const mockExamsRequest = async () => mockExamsResponse;

export const mockScheduleRequest = async () => mockScheduleResponse;

export const mockTranscriptRequest = async () => mockTranscriptResponse;
