import {useState} from "react"


export function useDoubleClick(click: () => void, doubleClick: () => void, delay = 300) {
    const [time, setTime] = useState<number | null>(null)

    function handleClick() {
        if (!time) {
            //@ts-ignore
            setTime(setTimeout(() => {
                click()
                setTime(null)
            }, delay))
        }
    }

    function handleDoubleClick() {
        if (time) clearTimeout(time)
        setTime(null)
        doubleClick()
    }

    return {
        onClick: handleClick,
        onDoubleClick: handleDoubleClick
    }
}