//
//  MapView.swift
//  ysmz-ios
//
//  Created by Yeonung Choi on 2022/11/19.
//

import SwiftUI
import MapKit
import CoreLocation

struct MapView: UIViewRepresentable {
    @StateObject var locationManager = LocationManager()

    func makeUIView(context: Context) -> MKMapView {
        MKMapView(frame: .zero)
    }

    var userCoordinate: (Double, Double) {
        return (
            // The default location is Yeoksam station.
            locationManager.lastLocation?.coordinate.latitude ?? 37.500723072486,
            locationManager.lastLocation?.coordinate.longitude ?? 127.03680544372
        )
    }

    func updateUIView(_ view: MKMapView, context: Context) {
        let coordinate = CLLocationCoordinate2D(
            latitude: userCoordinate.0, longitude: userCoordinate.1)
        let span = MKCoordinateSpan(latitudeDelta: 0.01, longitudeDelta: 0.01)
        let region = MKCoordinateRegion(center: coordinate, span: span)
        view.setRegion(region, animated: true)
    }
}

struct MapView_Previews: PreviewProvider {
    static var previews: some View {
        MapView()
    }
}
