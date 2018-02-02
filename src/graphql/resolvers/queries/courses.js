import R from 'ramda';
import {
  requestCoursework,
  requestAttendance,
  requestExams,
} from '../../../datasource/network';
import {
  parseCourses,
  parseCoursework,
  parseMidterms,
  parseAttendance,
  parseExams,
} from '../../../datasource/parser';
import { filterBy } from '../helpers';

const augmentCoursework = coursework => course =>
  R.assoc('coursework', R.filter(R.propEq('code', course.code))(coursework))(
    course,
  );

const augmentMidterm = midterms => course =>
  R.assoc('midterm', R.find(R.propEq('code', course.code))(midterms))(course);

const augmentAttendance = attendance => course =>
  R.assoc('absence', R.find(R.propEq('code', course.code))(attendance))(course);

const augmentExams = exams => course =>
  R.assoc('exam', R.find(R.propEq('code', course.code))(exams))(course);

const coursesResolver = async (obj, args, context) => {
  const { isAuthorized } = obj;
  const { code } = args;
  const { username, password } = context;

  if (!isAuthorized) {
    return null;
  }

  const [
    courseworkResponse,
    attendanceResponse,
    examsResponse,
  ] = await Promise.all([
    requestCoursework({ username, password }),
    requestAttendance({ username, password }),
    requestExams({ username, password }),
  ]);

  const courses = parseCourses(courseworkResponse);
  const coursework = parseCoursework(courseworkResponse);
  const midterms = parseMidterms(courseworkResponse);
  const attendance = parseAttendance(attendanceResponse);
  const exams = parseExams(examsResponse);

  return R.pipe(
    filterBy('code', code),
    R.map(augmentCoursework(coursework)),
    R.map(augmentMidterm(midterms)),
    R.map(augmentAttendance(attendance)),
    R.map(augmentExams(exams)),
  )(courses);
};

export default coursesResolver;
