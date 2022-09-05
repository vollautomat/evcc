export default {
  header: {
    docs: "Dokumentacija (Vokiečių k.)",
    blog: "Tinklaraštis",
    github: "GitHub",
    login: " Automobilių prisijungimai",
    about: "Apie EVCC",
    theme: {
      auto: "Design: System",
      light: "Design: Light",
      dark: "Design: Dark",
    },
  },
  footer: {
    version: {
      availableShort: "Naujinimas",
      availableLong: "Yra naujesnė versija",
      modalTitle: "Yra naujesnė versija",
      modalUpdateStarted: "Pasibaigus naujinimui EVCC startuos iš naujo..",
      modalInstalledVersion: "Dabartinė versija",
      modalNoReleaseNotes: 
      "Naujinimo detalių nėra. Daugiau informacijos rasite čia:",
      modalCancel: "Atšaukti",
      modalUpdate: "Naujinti",
      modalUpdateNow: "Naujinti dabar",
      modalDownload: "Atsisiųsti",
      modalUpdateStatusStart: "Naujinimas prasidėjo: ",
      modalUpdateStatusFailed: "Naujinimas nepavyko: ",
    },
    savings: {
      footerShort: "{percent}% saulės",
      footerLong: "{percent}% saulės energija",
      modalTitle: "Įkrovimo energijos apžvalga",
      sinceServerStart: "Nuo evcc starto {since}.",
      percentTitle: "Saulės energija",
      percentSelf: "{self} kWh saulės",
      percentGrid: "{grid} kWh tinklas",
      priceTitle: "Energijos kaina",
      priceFeedIn: "{feedInPrice} eksporto",
      priceGrid: "{gridPrice} tinklo",
      savingsTitle: "Sutaupyta",
      savingsComparedToGrid: "palyginus su tinklu",
      savingsTotalEnergy: "{total} kWh įkrauta",
    },
    sponsor: {
      thanks: "Ačiū, kad mus remiate {sponsor}! Taip prisidedate prie projekto vystymo.",
      confetti: "Norite konfeti?",
      supportUs:
        "Mūsų misija: Siekiame, kad įkrovimas saulės energija taptų standartu. Padėkite mums ir paremkite evcc finansiškai.",
      sticker: "... ar evcc lipdukų?",
      confettiPromise: "Gausite lipdukų ir skaitmeninių konfeti ;)",
      becomeSponsor: "Tapkite rėmėju",
    },
  },
  notifications: {
    modalTitle: "Pranešimai",
    dismissAll: "Išvalyti visus",
  },
  main: {
    energyflow: {
      noEnergy: "Nėra skaitiklių duomenų",
      homePower: "Namo suvartojimas",
      pvProduction: "Gamyba",
      loadpoints: "Įkroviklis | Įkroviklis | {count} Įkrovikliai",
      battery: "Baterija",
      batteryCharge: "Baterijos įkrovimas",
      batteryDischarge: "Baterijos iškrovimas",
      gridImport: "Tinklo importas",
      selfConsumption: "Sunaudojama iškart",
      pvExport: "Tinklo eksportas",
    },
    mode: {
      off: "Stop",
      minpv: "Min+PV",
      pv: "PV",
      now: "Greitas",
    },
    loadpoint: {
      fallbackName: "Įkroviklis",
      remoteDisabledSoft: "{source}: adaptyvus PV įkrovimas išjungtas",
      remoteDisabledHard: "{source}: išjungtas",
      power: "Galia",
      charged: "Įkrauta",
      duration: "Trukmė",
      remaining: "Liko",
    },
    loadpointSettings: {
      title: 'Nustatymai "{0}"',
      vehicle: "Automobilis",
      currents: "Įkraunama",
      minSoC: {
        label: "Minimali įkrova",
        description:
          'Minimali įkrova. Automobilis įkraunamas "greitai" iki {0}% PV nustatyme. Toliau įkraunamas tik saulės energijos pertekliumi.',
      },
      phasesConfigured: {
        label: "Fazės",
        phases_0: "automatinis perjungimas",
        phases_1: "1 fazė",
        phases_1_hint: "({min} to {max})",
        phases_3: "3 fazės",
        phases_3_hint: "({min} to {max})",
      },
      maxCurrent: {
        label: "Max. Srovė",
      },
      minCurrent: {
        label: "Min. Srovė",
      },
      default: "standartiškai",
      disclaimerHint: "Pastaba:",
      disclaimerText: "Šie pakeitimai neišlieka ir po EVCC serverio restarto pradings.",
    },
    vehicles: "Autoparkas",
    vehicle: {
      fallbackName: "Automobilis",
      vehicleSoC: "Įkrova",
      targetSoC: "Limitas",
      none: "Nėra automobilio",
      unknown: "Nežinomas automobilis",
      changeVehicle: "Pakeisti automobilį",
      detectionActive: "Bandome atpažinti automobilį ...",
    },
    vehicleSoC: {
      disconnected: "neprijungtas",
      charging: "vyksta įkrovimas",
      ready: "leidžiama įkrauti",
      connected: "automobilis prijungtas",
      vehicleTarget: "Automobilio limitas: {soc}%",
    },
    vehicleStatus: {
      minCharge: "minimalus įkrovimas iki {soc}%.",
      waitForVehicle: "Įkrovimas leidžiamas. Laukiama automobilio signalo.",
      vehicleTargetReached: "Automobilio limitas {soc}% pasiektas.",
      charging: "Įkraunama.",
      targetChargePlanned: "Suplanuotas įkrovimas, prasidės {time}.",
      targetChargeWaitForVehicle: "Suplanuotas įkrovimas leidžiamas. Laukiama automobilio signalo.",
      targetChargeActive: "Suplanuotas įkrovimas aktyvuotas.",
      connected: "Prijungtas.",
      pvDisable: "Trūksta saulės, įkrovimo pauzė už {remaining}.",
      pvEnable: "Saulės užtenka, įkrovimas prasidės už {remaining}.",
      scale1p: "Sumažinti į vienfazį įkrovimą už {remaining}.",
      scale3p: "Padidinti į trifazį įkrovimą už {remaining}.",
      disconnected: "Neprijungtas.",
      unknown: "",
    },
    provider: {
      login: "prisijungti",
      logout: "atsijungti",
    },
    targetCharge: {
      title: "Įkrauti iki",
      inactiveLabel: "Įkrauti iki",
      activeLabel: "{time}",
      modalTitle: "Nustatyti įkrovimo pabaigos laiką",
      setTargetTime: "nenustatytas",
      description: "Kada automobilis turėtų būti įkrautas iki {targetSoC}%?",
      today: "šiandien",
      tomorrow: "rytoj",
      targetIsInThePast: "Pasirinktas laikas yra praeityje.",
      remove: "Panaikinti",
      activate: "Aktyvuoti",
      experimentalLabel: "Eksperimentinis",
      experimentalText: `
        Ši funkcija veikia, bet dar nėra tobula. 
        Apie netikėtą elgesį praneškite mūsų
      `,
    },
  },
  offline: {
    message: "Nėra ryšio su serveriu.",
    reload: "Perkrauti?",
  },
};
