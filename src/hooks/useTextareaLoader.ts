import {MutableRefObject} from "react"


export default function useTextareaLoader(ref: MutableRefObject<HTMLTextAreaElement>, loading: boolean) {
    const height = ref.current.getBoundingClientRect().height

    if (loading) {

    }

}