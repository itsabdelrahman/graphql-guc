import { capitalize } from '../utilities';

describe('capitalize()', () => {
  test('It should capitalize the first character in a string', () => {
    expect(capitalize('quiz')).toBe('Quiz');
  });
});
