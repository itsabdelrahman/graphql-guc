import R from 'ramda';
import {
  requestLogin,
  requestAttendance,
  requestCoursework,
  requestExams,
  requestSchedule,
  requestTranscript,
} from './network';
import {
  parseLogin,
  parseAttendance,
  parseCourses,
  parseCoursework,
  parseMidterms,
  parseExams,
  parseSchedule,
  parseTranscript,
} from './parser';

export const isAuthorized = R.pipeP(requestLogin, parseLogin);

export const getAttendance = R.pipeP(requestAttendance, parseAttendance);

export const getCourses = R.pipeP(requestCoursework, parseCourses);

export const getCoursework = R.pipeP(requestCoursework, parseCoursework);

export const getMidterms = R.pipeP(requestCoursework, parseMidterms);

export const getExams = R.pipeP(requestExams, parseExams);

export const getSchedule = R.pipeP(requestSchedule, parseSchedule);

export const getTranscript = R.pipeP(requestTranscript, parseTranscript);
