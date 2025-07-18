export function removeIfExists<T>(array: Array<T>, checker: (e: T) => boolean) {
  const index = array.findIndex(checker);
  if (index >= 0) {
    array.splice(index, 1);
    return true;
  }
  return false;
}

export function getMapFromArray<T, K, V = T>(
  array: T[],
  keyGetter: (entity: T) => K,
  valGetter: (entity: T) => V = (e) => e as unknown as V,
): Map<K, V> {
  const map = new Map<K, V>();
  for (const entity of array) {
    map.set(keyGetter(entity), valGetter(entity));
  }
  return map;
}

export function createBatches<T>(array: T[], batchSize: number): T[][] {
  const batches: T[][] = [];
  for (let i = 0; i < array.length; i += batchSize) {
    batches.push(array.slice(i, i + batchSize));
  }
  return batches;
}

export function arrayUnorderedEquals<T>(src: T[], tar: T[]) {
  const srcSet = new Set<T>(src);
  const tarSet = new Set<T>(tar);
  if (srcSet.size !== tarSet.size) return false;
  return tar.every((t) => srcSet.has(t));
}

export function arrayOrderedEquals<T>(src: T[], tar: T[]) {
  if (src.length !== tar.length) return false;
  for (let i = 0; i < src.length; i += 1) {
    if (src[i] !== tar[i]) return false;
  }
  return true;
}

/**
 * Returns values in tar that are missing in src.
 */
export function getMissingValues<T>(src: T[], tar: T[]) {
  return tar.filter((v) => !src.includes(v));
}

export function dedupe<T, K>(array: T[], keyGetter: (entry: T) => K) {
  const seen = new Set<K>();
  return array.filter((entry) => {
    const key = keyGetter(entry);
    if (seen.has(key)) return false;
    seen.add(key);
    return true;
  });
}
