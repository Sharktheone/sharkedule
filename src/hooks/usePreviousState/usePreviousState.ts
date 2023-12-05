import {useState} from "react"


export function usePreviousState<S = undefined>(initial: S) {
    const [state, setState] = useState<S>(initial)
    const [previousState, setPreviousState] = useState<S | undefined>(undefined)

    function updateState(newState: S) {
        setPreviousState(state)
        setState(newState)
    }

    return [state, updateState, previousState] as const
}