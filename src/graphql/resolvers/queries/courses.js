import R from 'ramda';
import { requestCoursework } from '../../../datasource/network';
import {
  parseCourses,
  parseCoursework,
  parseMidterms,
} from '../../../datasource/parser';
import { filterBy } from '../helpers';

const augmentCoursework = coursework => course =>
  R.assoc('coursework', R.filter(R.propEq('code', course.code))(coursework))(
    course,
  );

const augmentMidterm = midterms => course =>
  R.assoc('midterm', R.find(R.propEq('code', course.code))(midterms))(course);

const coursesResolver = async (obj, args, context) => {
  const { code } = args;
  const { username, password } = context;

  const response = await requestCoursework({ username, password });

  const courses = parseCourses(response);
  const coursework = parseCoursework(response);
  const midterms = parseMidterms(response);

  return R.pipe(
    filterBy('code', code),
    R.map(augmentCoursework(coursework)),
    R.map(augmentMidterm(midterms)),
  )(courses);
};

export default coursesResolver;
