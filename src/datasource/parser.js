import R from 'ramda';
import moment from 'moment';
import { capitalize } from '../utilities';

const transformAttendance = element => ({
  code: R.replace(/\s/g, '', element.Code),
  name: R.trim(element.Name),
  level: Number(element.AbsenceLevel),
});

const transformCoursework = aggregation =>
  R.map(element => {
    const course = R.find(currentCourse =>
      R.equals(currentCourse.sm_crs_id, element.sm_crs_id),
    )(aggregation.CurrentCourses);
    return {
      code: R.pipe(R.match(/\((.*?)\)/), R.view(R.lensIndex(1)))(
        course.course_short_name,
      ),
      name: R.pipe(R.replace(/\((.*?)\)/, ''), R.trim)(
        course.course_short_name,
      ),
      type: R.pipe(R.trim, capitalize)(element.eval_method_name),
      grade: Number(element.grade),
      maximumGrade: Number(element.max_point),
    };
  })(aggregation.CourseWork);

const transformMidterms = element => ({
  code: R.pipe(
    R.split('-'),
    R.view(R.lensIndex(1)),
    R.trim,
    R.split(' '),
    R.last,
    R.trim,
  )(element.course_full_name),
  name: R.pipe(
    R.split('-'),
    R.view(R.lensIndex(1)),
    R.trim,
    R.split(' '),
    R.dropLast(1),
    R.join(' '),
    R.trim,
  )(element.course_full_name),
  grade: Number(element.total_perc),
});

const transformExams = element => ({
  code: R.pipe(
    R.replace(/^.+-/, ''),
    R.trim,
    R.split(' '),
    R.take(1),
    R.join(' '),
  )(element.course_name),
  name: R.pipe(
    R.replace(/^.+-/, ''),
    R.trim,
    R.split(' '),
    R.drop(1),
    R.join(' '),
  )(element.course_name),
  venue: element.rsrc_code,
  seat: element.seat_code,
  startsAt: moment(
    R.replace(/\s\s/g, ' ', element.start_time),
    'MMM D YYYY h:mmA',
    true,
  ).toISOString(),
});

const transformSchedule = element => ({
  code: R.replace(/\s/g, '', element.course_short_code),
  name: R.pipe(R.split('-'), R.take(1), R.join(' '), R.trim)(element.course),
  type: R.equals('Tut', element.class_type)
    ? 'TUTORIAL'
    : R.toUpper(element.class_type),
  weekday: R.toUpper(element.weekday),
  slot: Number(element.scd_col),
  venue: R.trim(element.location),
});

export const parseLogin = response => R.pathEq(['data', 'd'], 'True')(response);

export const parseAttendance = response =>
  R.pipe(
    R.path(['data', 'd']),
    JSON.parse,
    R.prop('AbsenceReport'),
    R.map(transformAttendance),
  )(response);

export const parseCoursework = response =>
  R.pipe(R.path(['data', 'd']), JSON.parse, transformCoursework)(response);

export const parseMidterms = response =>
  R.pipe(
    R.path(['data', 'd']),
    JSON.parse,
    R.prop('Midterm'),
    R.map(transformMidterms),
  )(response);

export const parseExams = response =>
  R.pipe(R.path(['data', 'd']), JSON.parse, R.map(transformExams))(response);

export const parseSchedule = response =>
  R.pipe(R.path(['data', 'd']), JSON.parse, R.map(transformSchedule))(response);
