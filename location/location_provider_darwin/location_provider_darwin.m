#import <CoreLocation/CoreLocation.h>
#import <Foundation/Foundation.h>
#import "location_provider_darwin.h"

@interface LocationManager : NSObject <CLLocationManagerDelegate>
{
    CLLocationManager *manager;
}

@property (readonly) NSInteger errorCode;

- (CLLocation *)getCurrentLocation;

@end

@implementation LocationManager

- (id)init {
    self = [super init];
    if (self) {
        manager = [[CLLocationManager alloc] init];
        manager.delegate = self;
        manager.desiredAccuracy = kCLLocationAccuracyBest;

        CLAuthorizationStatus status = manager.authorizationStatus;
        
        // Check authorization status and request if not determined
        if (status == kCLAuthorizationStatusNotDetermined) {
            [manager requestAlwaysAuthorization]; // macOS generally uses always authorization
            CFRunLoopRun();
        } else if (status == kCLAuthorizationStatusAuthorized || status == kCLAuthorizationStatusAuthorizedAlways) {
            [manager requestLocation]; // Proceed if already authorized
        } else {
            NSLog(@"Location authorization denied or restricted.");
        }
    }
    return self;
}

- (void)dealloc {
    [manager release];
    [super dealloc];
}

- (CLLocation *)getCurrentLocation {
    // Request location update
    [manager requestLocation];

    // Run the run loop until a location update or error occurs
    CFRunLoopRun();

    // If there's an error, return nil
    if (self.errorCode != 0) {
        return nil;
    }

    CLLocation *location = manager.location;

    // Ensure valid location data
    if (location.horizontalAccuracy < 0.0) {
        return nil;
    }

    return location;
}

- (void)locationManager:(CLLocationManager *)manager didUpdateLocations:(NSArray<CLLocation *> *)locations {
    // Stop the run loop when a location update is received
    CFRunLoopStop(CFRunLoopGetCurrent());
}

- (void)locationManager:(CLLocationManager *)manager didFailWithError:(NSError *)error {
    _errorCode = error.code; // Capture error code
    CFRunLoopStop(CFRunLoopGetCurrent()); // Stop the run loop on error
}

- (void)locationManager:(CLLocationManager *)locationManager didChangeAuthorizationStatus:(CLAuthorizationStatus)status {
    if (status == kCLAuthorizationStatusAuthorized || status == kCLAuthorizationStatusAuthorizedAlways) {
        [manager requestLocation];
        CFRunLoopStop(CFRunLoopGetCurrent());
    } else if (status == kCLAuthorizationStatusDenied || status == kCLAuthorizationStatusRestricted) {
        NSLog(@"Location authorization denied or restricted.");
        CFRunLoopStop(CFRunLoopGetCurrent());
    }

}

@end

int get_current_location(Location *loc) {
    if (![CLLocationManager locationServicesEnabled]) {
        NSLog(@"Location services disabled");
        return kCLErrorLocationUnknown;
    }

    @autoreleasepool {
        LocationManager *locationManager = [[LocationManager alloc] init];
        CLLocation *clloc = [locationManager getCurrentLocation];

        if (locationManager.errorCode != 0) {
            return locationManager.errorCode;
        }

        // Populate the location struct with obtained data
        loc->coordinate = clloc.coordinate;
        loc->altitude = clloc.altitude;
        loc->horizontalAccuracy = clloc.horizontalAccuracy;
        loc->verticalAccuracy = clloc.verticalAccuracy;
    }

    return 0;
}
