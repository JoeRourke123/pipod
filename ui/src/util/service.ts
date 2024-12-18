export const fetchAuthStatus = async () => {
    const response = await fetch(
        `http://localhost:9091/isAuth`
    );

    return await response.json()
}

export const fetchSpotifyToken = async () => {
    const responseJson = await fetchAuthStatus();

    return responseJson["accessToken"];
}

export const player = async (playerData: any) => {
    return await fetch("http://localhost:9091/player", {
        method: "POST",
        body: JSON.stringify(playerData),
        headers: {
            "Content-Type": "application/json"
        }
    });
}