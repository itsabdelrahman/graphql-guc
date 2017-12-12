import R from 'ramda';
import moment from 'moment';

const transformAttendance = element => ({
  code: R.trim(element.Code),
  name: R.trim(element.Name),
  level: Number(element.AbsenceLevel),
});

const transformExamsSchedule = element => ({
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

export const parseLogin = response => R.pathEq(['data', 'd'], 'True')(response);

/** @TODO */
export const parseCoursework = response => ({});

export const parseAttendance = response =>
  R.pipe(
    R.path(['data', 'd']),
    JSON.parse,
    R.prop('AbsenceReport'),
    R.map(transformAttendance),
  )(response);

export const parseExamsSchedule = response =>
  R.pipe(R.path(['data', 'd']), JSON.parse, R.map(transformExamsSchedule))(
    response,
  );

/** @TODO */
export const parseSchedule = response => ({});
