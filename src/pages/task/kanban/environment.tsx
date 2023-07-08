import {createContext, Dispatch, ReactNode, SetStateAction} from "react"
import {environment} from "@kanban/types"

type context = {
    environment: environment
    setEnvironment: Dispatch<SetStateAction<environment>>
}

export const EnvironmentContext = createContext<context>({} as context)


type Props = {
    children: ReactNode
    environment: environment
    setEnvironment: Dispatch<SetStateAction<environment>>
}

export function EnvironmentProvider({children, setEnvironment, environment}: Props) {
    const value: context = {
        environment,
        setEnvironment
    }
    return (
        <EnvironmentContext.Provider value={value}>
            {children}
        </EnvironmentContext.Provider>
    )
}