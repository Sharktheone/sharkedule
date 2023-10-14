import defLogo from '../assets/default.png'

export function logo() {
    const defaultLogo = true //TODO Get from config context

    if (defaultLogo) {
        return defLogo
    } else {
        //TODO Get from config context
        return defLogo
    }
}