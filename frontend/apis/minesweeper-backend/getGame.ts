const getGame = async (gameID: string): Promise<any> => {
  // TODO: think about replacing with config file

  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_HOST}/api/games/${gameID}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        cache: "no-store",
      }
    );

    if (!response.ok || response.status !== 200) {
      return Promise.reject("Error getting game");
    }

    // ! TODO: add game model and converter
    return response.json();
  } catch (error) {
    throw error;
  }
};

export default getGame;
