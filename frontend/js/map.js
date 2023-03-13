let map;
let placemark;
const defaultCenter = [55.76, 37.64];
const defaultZoom = 11;

const deliveryButton = document.getElementById("deliveryButton");
const mapError = document.getElementById("map-error")


ymaps.ready(init);

function init() {
    map = new ymaps.Map("map", {
        center: defaultCenter, // Москва
        zoom: defaultZoom
    }, {
        searchControlProvider: 'yandex#search'
    });
}

deliveryButton.addEventListener("click", () => {
    const addr = document.getElementById("deliveryAddress").value;
    console.log("адрес = ", addr);

    if (addr === "") {
        showError("Введен пустой адрес")
    } else {
        geocode(addr);
    }
})

function showError(message) {
    mapError.textContent = message;
    map.setCenter(defaultCenter, defaultZoom);
}

function geocode(addr) {
    ymaps.geocode(addr).then(function (res) {
        let obj = res.geoObjects.get(0),
            error;

        if (obj) {
            switch (obj.properties.get('metaDataProperty.GeocoderMetaData.precision')) {
                case 'exact':
                    break;
                case 'number':
                case 'near':
                case 'range':
                    error = 'Уточните номер дома';
                    break;
                case 'street':
                    error = 'Уточните номер дома';
                    break;
                case 'other':
                default:
                    error = 'Неточный адрес, требуется уточнение';
            }
        } else {
            error = 'Адрес не найден';
        }

        if (error) {
            showError(error);
        } else {
            showResult(obj);
        }
    }, function (e) {
        console.log("error = ", e)
    })
}

function showResult(obj) {
    mapError.textContent = "Адрес подтвержден";
    mapError.classList.add("note")

    let mapState = ymaps.util.bounds.getCenterAndZoom(
        obj.properties.get('boundedBy'),
        [100, 100]
    );

    let shortAddress = [obj.getThoroughfare(), obj.getPremiseNumber(), obj.getPremise()].join(' ');

    mapState.controls = [];
    createMap(mapState, shortAddress);
}

function createMap(state, caption) {
    if (!placemark) {
        map.setCenter(state.center, state.zoom);
        placemark = new ymaps.Placemark(
            map.getCenter(), {
                iconCaption: caption,
                balloonContent: caption
            }, {
                preset: 'islands#redDotIconWithCaption'
            });
        map.geoObjects.add(placemark);
    } else {
        map.setCenter(state.center, state.zoom);
        placemark.geometry.setCoordinates(state.center);
        placemark.properties.set({iconCaption: caption, balloonContent: caption});
    }
}
