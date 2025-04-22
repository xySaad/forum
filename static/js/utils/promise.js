export const resolve = async (promise) => {
    let result, err = null
    try {
        result = await promise
    } catch (error) {
        err = error
    }

    return [result, err]
}