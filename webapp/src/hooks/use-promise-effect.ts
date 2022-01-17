import {DependencyList, useEffect, useState} from 'react'


export function usePromiseEffect<T>(effect: () => Promise<T>, deps: DependencyList) {
    const [value, setValue] = useState<T>(null)

    useEffect(() => {
        effect()
            .then((value) => setValue(value))
            .catch((error) => {
                setValue(value)
                throw(error);
            })
    }, deps)

    return value
}