/**
 * Return true is obj is and object.
 * Will return false if obj is null or is an Array
 * @param obj
 * @returns
 */
export const IsObject = (obj: any): boolean =>
  typeof obj === "object" && !Array.isArray(obj) && obj !== null;

/**
 * ArrayToObject will take and array of object and
 * @param array
 * @param property
 * @returns
 */
export const ArrayToObject = <T>(
  array: Array<T>,
  property: string = "id"
): Record<string, T> => {
  console.log(array, property)
  return array.reduce((acc, obj) => {
    if (!IsObject(obj)) {
      throw new Error(`Not an object, found: ${typeof obj}`);
    }
    const key = obj[property];
    if (!key) {
      throw new Error(
        `Invalid key: ${key} -- ${key} is missing in object: ${JSON.stringify(
          obj
        )}`
      );
    }
    return { ...acc, [key]: { ...obj } };
  }, {});}
