export const stringToColor = (string: string): string => {
  let hash = 0;
  let i;

  /* eslint-disable no-bitwise */
  for (i = 0; i < string.length; i += 1) {
    hash = string.charCodeAt(i) + ((hash << 5) - hash);
  }

  let color = "#";

  for (i = 0; i < 3; i += 1) {
    const value = (hash >> (i * 8)) & 0xff;
    color += `00${value.toString(16)}`.substr(-2);
  }
  /* eslint-enable no-bitwise */

  return color;
};

export const getInitials = (name = ""): string =>
  name
    .replace(/\s+/, " ")
    .split(" ")
    .slice(0, 2)
    .map((v) => v && v[0].toUpperCase())
    .join("");


export const labelize = (name: string): string => {
  const acronyms: string[] = ['id', 'url', 'uri', 'json', 'api']
  let result = name.toLowerCase().replace("_", " ")
  acronyms.map((acronym) => { result = result.replace(acronym, acronym.toUpperCase())})
  return result.replace(/(^|\s)\S/g, function(t) { return t.toUpperCase() });
}
