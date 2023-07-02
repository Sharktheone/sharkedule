import {MenuView} from "@/components/menu/types"

type Props = {
    view: MenuView
}

export default function RenderView({view}: Props) {
    console.log(typeof view.items)
    return (
        <div>
            {view?.name}
        </div>
    )

}